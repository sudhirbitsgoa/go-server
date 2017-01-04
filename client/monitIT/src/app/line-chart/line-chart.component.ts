import { Component, OnInit } from '@angular/core';
import {MetricServiceService} from '../metric-service.service'

@Component({
  selector: 'app-line',
  templateUrl: './line-chart.component.html',
  styleUrls: ['./line-chart.component.css'],
  providers: [MetricServiceService]
})
export class LineChartComponent implements OnInit {
  private metricService: MetricServiceService;
  constructor(metricService: MetricServiceService) {
      this.metricService = metricService;
  }
  // lineChart
  public lineChartData: Array<any> = [
    { data: [65, 59, 80, 81, 56, 55, 40], label: 'Series A' },
    { data: [28, 48, 40, 19, 86, 27, 90], label: 'Series B' },
    { data: [18, 48, 77, 9, 100, 27, 40], label: 'Series C' }
  ];
  public lineChartLabels: Array<any> = ['January', 'February', 'March', 'April', 'May', 'June', 'July'];
  public lineChartOptions: any = {
    animation: false,
    responsive: true
  };
  ngOnInit() {
    this.metricService.GetCPUMEMMetric(data => {
        console.log('this is data', data);
    });
    // (function(self) {
    //   var conn;
    //   // var msg = $("#msg");
    //   // var log = $("#log");
    //   if (window["WebSocket"]) {
    //     conn = new WebSocket("ws://localhost:8080/ws");
    //     conn.onopen = function () {
    //         conn.send('cpu');
    //     }
    //     conn.onclose = function(evt) {
    //       // appendLog($("<div><b>Connection closed.</b></div>"))
    //     }
    //     conn.onmessage = function(evt) {
    //       console.log(evt.data)
    //       let data = JSON.parse(evt.data);
    //       self.lineChartData[0].data.push(data.perct);
    //
    //       // appendLog($("<pre/>").text(evt.data))
    //     }
    //   } else {
    //     // appendLog($("<div><b>Your browser does not support WebSockets.</b></div>"))
    //   }
    // })(this);
  }
  public lineChartColors: Array<any> = [
    { // grey
      backgroundColor: 'rgba(148,159,177,0.2)',
      borderColor: 'rgba(148,159,177,1)',
      pointBackgroundColor: 'rgba(148,159,177,1)',
      pointBorderColor: '#fff',
      pointHoverBackgroundColor: '#fff',
      pointHoverBorderColor: 'rgba(148,159,177,0.8)'
    },
    { // dark grey
      backgroundColor: 'rgba(77,83,96,0.2)',
      borderColor: 'rgba(77,83,96,1)',
      pointBackgroundColor: 'rgba(77,83,96,1)',
      pointBorderColor: '#fff',
      pointHoverBackgroundColor: '#fff',
      pointHoverBorderColor: 'rgba(77,83,96,1)'
    },
    { // grey
      backgroundColor: 'rgba(148,159,177,0.2)',
      borderColor: 'rgba(148,159,177,1)',
      pointBackgroundColor: 'rgba(148,159,177,1)',
      pointBorderColor: '#fff',
      pointHoverBackgroundColor: '#fff',
      pointHoverBorderColor: 'rgba(148,159,177,0.8)'
    }
  ];
  public lineChartLegend: boolean = true;
  public lineChartType: string = 'line';

  public randomize(): void {

    let _lineChartData: Array<any> = new Array(this.lineChartData.length);
    for (let i = 0; i < this.lineChartData.length; i++) {
      _lineChartData[i] = { data: new Array(this.lineChartData[i].data.length), label: this.lineChartData[i].label };
      for (let j = 0; j < this.lineChartData[i].data.length; j++) {
        _lineChartData[i].data[j] = Math.floor((Math.random() * 100) + 1);
      }
    }
  }

  // events
  public chartClicked(e: any): void {
    console.log(e);
  }

  public chartHovered(e: any): void {
    console.log(e);
  }
}
