import http from 'k6/http';
import { check, sleep } from 'k6';
import exec from 'k6/execution';
import { textSummary } from 'https://jslib.k6.io/k6-summary/0.0.1/index.js';

export const options = {
  stages: [
    { duration: '10s', target: 1000 },
    { duration: '10s', target: 5000 },
    { duration: '30s', target: 10000 },
  ],
};

export function setup() {
  return http.response("https://lemur-0.cloud-iam.com/auth/realms/keycloak-overflow/protocol/openid-connect/token", 
  { client_id: "client-overflow", grant_type: "password", username: "k6@stevenjansen.com", password: "K6Password" })
  .json("access_token")
}

export default function (accessToken) {
  const res = http.get('http://steven.krijgtochdetyf.us/responses/api/response/getAll', { headers: { 'Authorization': "Bearer " + accessToken } });
  // if(!check(res, { 'status must be 200': (r) => r.status == 200 })){
  //   exec.test.abort()
  // }
  sleep(500);
}


export function handleSummary(data) {
  console.log('Finished executing performance tests');

  return {
    'stdout': textSummary(data, { indent: ' ', enableColors: true }), // Show the text summary to stdout...
    'summary.txt': textSummary(data, { indent: ' ' }), // and a JSON with all the details...
  };
}