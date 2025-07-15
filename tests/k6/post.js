import http from 'k6/http';
import { uuidv4 } from 'https://jslib.k6.io/k6-utils/1.4.0/index.js';
import { check, sleep} from 'k6';

export const options = {
    vus: 100,    // кол-во пользователей
    duration: '30s', // длительонсть теста
};

export default function () {
    const url = 'http://host.docker.internal:8080/notifications';

    const payload = JSON.stringify({
        user_id: uuidv4(),
        message: `Test message`
    });
    const params = {
        headers: {
            'Content-Type': 'application/json',
        },
    };

    const res = http.post(url, payload, params);

    check(res, {
        'status is 201': (r) => r.status === 201,
    });
    sleep(1);
}