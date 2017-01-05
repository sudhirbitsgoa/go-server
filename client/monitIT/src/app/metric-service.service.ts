import { Injectable } from '@angular/core';
// import {$WebSocket, WebSocketSendMode} from 'angular2-websocket/angular2-websocket'

@Injectable()
export class MetricServiceService {
  private ws: WebSocket;
  constructor() {
      this.ws = new WebSocket("ws://localhost:8080/ws");
  }
  GetCPUMEMMetric(cb) {
    // this.ws.onMessage((evt: MessageEvent) => {
    //     cb(evt.data);
    // })
    // this.ws.onMessageHandler((evt: MessageEvent) => {
    //     console.log('whats wrong')
    // })
    // this.ws.send('cpu')//.subscribe(
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
    // this.ws.send('cpu');
    this.ws.onopen = () => {
        this.ws.send('cpu');
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
