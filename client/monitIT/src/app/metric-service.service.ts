import { Injectable } from '@angular/core';
// import {$WebSocket, WebSocketSendMode} from 'angular2-websocket/angular2-websocket'

@Injectable()
export class MetricServiceService {
  private ws: WebSocket;
  constructor() {
      this.ws = new WebSocket("ws://localhost:8080/ws");
  }
  GetCPUMEMMetric(cb) {
    this.ws.onopen = () => {
        this.ws.send('cpu');
        this.ws.send('memory');
    }
    this.ws.onclose = function(evt) {
      // appendLog($("<div><b>Connection closed.</b></div>"))
    }
    this.ws.onmessage = function(evt) {
      console.log(evt.data);
      let data = JSON.parse(evt.data);
      cb(data);
    }
  }
}
