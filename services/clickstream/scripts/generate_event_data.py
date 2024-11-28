#!/usr/bin/python
import csv
import sys
import random
from typing import Optional
from faker import Faker
from datetime import datetime
from dataclasses import dataclass, field, asdict

f = Faker()
f.seed_instance(1)


@dataclass
class Session:
    id: str = field(default_factory=lambda: str(f.uuid4()))

    def as_dict(self):
        return asdict(self)


@dataclass
class Event:
    sid: str
    productId: Optional[str] = None
    searchQuery: Optional[str] = None
    channel: str = field(default_factory=lambda: random.choice(["web", "mobile"]))
    action: str = field(default_factory=lambda: random.choice(["click", "impression"]))
    type: str = field(default_factory=lambda: random.choice(["search", "component"]))
    timestamp: str = field(
        default_factory=lambda: str(
            datetime.isoformat(
                f.date_time_between(
                    start_date=datetime.fromisoformat("2024-01-01T00:00:00"),
                    end_date=datetime.fromisoformat("2024-01-01T23:59:59"),
                )
            )
        )
    )
    culture: str = field(
        default_factory=lambda: random.choice(["tr_TR", "en_UK", "en_US", "en_AU"])
    )
    ipAddress: str = field(default_factory=lambda: str(f.ipv4()))
    latitude: float = field(default_factory=lambda: float(f.latitude()))
    longitude: float = field(default_factory=lambda: float(f.longitude()))

    def as_dict(self):
        return asdict(self)


if __name__ == "__main__":
    args = sys.argv

    total_session = 100
    if len(args) > 1:
        total_session = int(args[1])

    sessions = [Session() for _ in range(1, total_session + 1)]
    products = [f.random_int(0, 1000) for _ in range(1, 10)]

    random_events = [
        Event(
            sid=session.id,
            productId=str(random.choice(products)),
        ).as_dict()
        for session in sessions
    ]

    sorted_events = sorted(
        random_events, key=lambda x: datetime.fromisoformat(x["timestamp"])
    )

    def transform(e: dict):
        if e.get("type") == "search" and e.get("searchQuery") is None:
            e["searchQuery"] = f.color_name()
            e["productId"] = None

        return e

    transformed_events = list(map(transform, sorted_events))

    csv_file_path = f"random_data_{total_session}.csv"
    with open(csv_file_path, "w") as csv_file:
        writer = csv.DictWriter(csv_file, fieldnames=transformed_events[0].keys())
        writer.writeheader()
        writer.writerows(transformed_events)

    print(f"Random events saved into {csv_file_path}")
