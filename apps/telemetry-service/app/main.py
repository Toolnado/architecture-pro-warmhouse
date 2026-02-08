from datetime import datetime
from typing import List

from fastapi import FastAPI
from pydantic import BaseModel

from .database import get_connection, init_db

app = FastAPI(title="Telemetry Service", version="1.0.0")

init_db()


class TelemetryCreate(BaseModel):
    device_id: str
    metric_type: str
    value: float
    unit: str


class TelemetryRecord(BaseModel):
    id: int
    device_id: str
    metric_type: str
    value: float
    unit: str
    timestamp: datetime


@app.get("/health")
def health():
    return {"status": "Telemetry Service is running"}


@app.post("/api/telemetry", response_model=TelemetryRecord, status_code=201)
def create_telemetry(data: TelemetryCreate):
    conn = get_connection()
    cur = conn.cursor()
    cur.execute(
        """INSERT INTO telemetry (device_id, metric_type, value, unit)
           VALUES (%s, %s, %s, %s)
           RETURNING id, device_id, metric_type, value, unit, timestamp""",
        (data.device_id, data.metric_type, data.value, data.unit),
    )
    record = cur.fetchone()
    conn.commit()
    cur.close()
    conn.close()
    return record


@app.get("/api/telemetry/{device_id}", response_model=List[TelemetryRecord])
def get_telemetry(device_id: str):
    conn = get_connection()
    cur = conn.cursor()
    cur.execute(
        """SELECT id, device_id, metric_type, value, unit, timestamp
           FROM telemetry WHERE device_id = %s
           ORDER BY timestamp DESC LIMIT 100""",
        (device_id,),
    )
    records = cur.fetchall()
    cur.close()
    conn.close()
    return records
