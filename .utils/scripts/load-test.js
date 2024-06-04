import { check, group } from 'k6';
import http from 'k6/http';

export const options = {
   stages: [
       { duration: '0.2m', target: 3 }, // simulate ramp-up of traffic from 1 to 3 virtual users over 0.5 minutes.
       { duration: '0.2m', target: 4}, // stay at 4 virtual users for 0.5 minutes
       { duration: '0.2m', target: 0 }, // ramp-down to 0 users
     ],
};

export default function () {
   group('API uptime heath check', () => {
       const response = http.get('http://localhost:8080/health');
       check(response, {
           "status code should be 200": res => res.status === 200,
       });
   });

   group('API uptime GET attendants', () => {
       const response = http.get('http://localhost:8080/attendants');
       check(response, {
           "status code should be 200": res => res.status === 200,
       });
   });

   group('API uptime GET customers', () => {
       const response = http.get('http://localhost:8080/customers');
       check(response, {
           "status code should be 200": res => res.status === 200,
       });
    });
};

