import { Component, OnInit, Input } from '@angular/core';
import {MetricServiceService} from '../metric-service.service'

@Component({
  selector: 'app-line',
  templateUrl: './line-chart.component.html',
  styleUrls: ['./line-chart.component.css'],
  providers: [MetricServiceService]
})
export class LineChartComponent implements OnInit {
  @Input() metric: string;
  private metricService: MetricServiceService;
  private grapData: Array<Number>
  private timeStamp: Array<any>
  constructor(metricService: MetricServiceService) {
      this.metricService = metricService;
  }
  // lineChart
  public lineChartData: Array<any> = [
    { data: [], label: 'Series A' },
    // { data: [28, 48, 40, 19, 86, 27, 90], label: 'Series B' },
    // { data: [18, 48, 77, 9, 100, 27, 40], label: 'Series C' }
  ];
  public lineChartLabels: Array<any> = ['January', 'February', 'March', 'April', 'May', 'June', 'July'];
  public lineChartOptions: any = {
    animation: false,
    responsive: true
  };
  ngOnInit() {
    this.grapData = []
    this.timeStamp = []
    this.metricService.GetCPUMEMMetric(data => {
        this.grapData.push(data.perct);
        let date = new Date(data.timeStamp);
        this.timeStamp.push(`${date.getHours()}:${date.getMinutes()}:${date.getSeconds()}`);
        // this.lineChartData[0].data.push(data.perct);
        if (data.metric === this.metric) {
            this.lineChartData =  [{ data: this.grapData, label: data.metric }];
        }
        this.lineChartLabels = this.timeStamp;
    });
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

  // events
  public chartClicked(e: any): void {
    console.log(e);
  }

  public chartHovered(e: any): void {
    console.log(e);
  }
}
