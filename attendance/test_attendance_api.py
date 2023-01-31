import pytest
import json
from attendance_api import app

def test_index_route():
    response = app.test_client().get('/')
    assert response.status_code == 404

@pytest.mark.get_request
def test_attendance_search():
  response = app.test_client().get('/attendance/search')
  res = json.loads(response.data.decode('utf-8')).get("message")
  assert res == "Error while pulling data for attendance"
  assert type(res) is str

@pytest.mark.get_health
def test_attendance_health():
  response = app.test_client().get('/attendance/healthz')
  res = json.loads(response.data.decode('utf-8')).get("description")
  assert res == "MySQL is not healthy"
  assert type(res) is str