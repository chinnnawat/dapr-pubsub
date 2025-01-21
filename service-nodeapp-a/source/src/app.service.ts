import { Injectable } from '@nestjs/common';
import { HttpService } from '@nestjs/axios';

@Injectable()
export class AppService {
  constructor(private readonly httpService: HttpService) {}
  current_time = new Date().toLocaleString();

  getHello(): string {
    return 'Hello World! From Service A';
  }

  invoke(): any {
    this.httpService
      .get('http://127.0.0.1:3500/v1.0/invoke/service-b/method/api')
      .subscribe((response) => {
        console.log(response.data);
      });

    return 'Check the logs for the response from Service B, C';
  }

  pubsub(): any {
    this.httpService
      .post(
        'http://localhost:3500/v1.0/publish/pubsub/message-to-service-b',
        {
          message: 'Hello from Service A Time: ' + this.current_time,
        },
        {
          headers: {
            'Content-Type': 'application/json',
          },
        },
      )
      .subscribe((response) => {
        console.log(response.status);
      });

    return 'Check the logs for the response from Service B, C';
  }
}
