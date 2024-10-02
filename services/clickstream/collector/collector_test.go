package collector

import (
	"bytes"
	"context"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	tc "github.com/testcontainers/testcontainers-go/modules/compose"
)

func Test(t *testing.T) {
	// given
	compose, err := tc.NewDockerCompose("test-kafka-docker-compose.yml")
	require.NoError(t, err, "NewDockerComposeAPI()")
	t.Cleanup(func() {
		require.NoError(t, compose.Down(context.Background(), tc.RemoveOrphans(true), tc.RemoveImagesLocal), "compose.Down()")
	})

	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	require.NoError(t, compose.Up(ctx, tc.Wait(true)), "compose.Up()")

	tLogger := log.New(os.Stdout, "test", log.Ltime)
	producer, err := NewProducer(
		tLogger,
		"localhost",
		"9093",
		"test-topic",
	)
	require.NoError(t, err, "NewProducer()")

	t.Run("handleEvent", func(t *testing.T) {
		testPayload := []byte(`{"tv001": "1", "tv002": "1", "tv003": "web"}`)
		req := httptest.NewRequest(http.MethodPost, "/e", bytes.NewBuffer(testPayload))
		w := httptest.NewRecorder()

		handleEvent(tLogger, nil, producer).ServeHTTP(w, req)
		res := w.Result()
		defer res.Body.Close()
		data, err := io.ReadAll(res.Body)

		require.NoError(t, err, "handleEvent()")
		require.Equal(t, int(0), len(data))
		require.Equal(t, http.StatusOK, res.StatusCode)
		require.Equal(t, int64(-1), res.ContentLength)
	})

	t.Run("handlePixel", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/pixel?tv001=1&tv002=1&tv003=web", nil)
		w := httptest.NewRecorder()

		handlePixel(tLogger, nil, producer).ServeHTTP(w, req)
		res := w.Result()
		defer res.Body.Close()
		data, err := io.ReadAll(res.Body)

		require.NoError(t, err, "handlePixel()")
		require.Equal(t, int(42), len(data))
		require.Equal(t, http.StatusOK, res.StatusCode)
		require.Equal(t, "image/gif", res.Header.Get("Content-Type"))
	})

}
