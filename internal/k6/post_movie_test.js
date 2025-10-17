import http from 'k6/http';
import { check } from 'k6';

export let options = {
  vus: 20,
  duration: '30s',
};

export default function () {
  const payload = JSON.stringify({
    title: `Test Movie`,
    year: 2024
  });

  const headers = { 'Content-Type': 'application/json' };

  let res = http.post('http://localhost:8080/movies', payload, { headers });

  check(res, {
    'status is 201': (r) => r.status === 201,
  });
}