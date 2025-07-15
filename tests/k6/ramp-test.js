import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
    stages: [
        { duration: '30s', target: 50},
        { duration: '30s', target: 100},
        { duration: '30s', target: 200},
        { duration: '30s', target: 300},
        { duration: '30s', target: 400},
        { duration: '30s', target: 500},
        { duration: '30s', target: 0},
    ],
};

export default function () {
    const res = http.get('http://host.docker.internal:8080/notifications');
    check(res, { 'status is 200': (r) => r.satus === 200 });
    sleep(1);
}