import { Injectable } from '@angular/core';
import {$WebSocket, WebSocketSendMode} from 'angular2-websocket/angular2-websocket'

@Injectable()
export class MetricServiceService {
  private ws: $WebSocket;
  constructor() {
      this.ws = new $WebSocket("ws://localhost:8080/ws");
  }
  GetCPUMEMMetric(cb) {
    this.ws.onMessage((evt: MessageEvent) => {
        cb(evt.data);
    })
    // this.ws.onMessageHandler((evt: MessageEvent) => {
    //     console.log('whats wrong')
    // })
    this.ws.send('cpu')//.subscribe(
    //     (msg)=> {
    //         console.log("next", msg.data);
    //         cb(msg.data);
    //     },
    //     (msg)=> {
    //         console.log("error", msg);
    //     },
    //     ()=> {
    //         console.log("complete");
    //     })
  }
}
