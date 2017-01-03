/* tslint:disable:no-unused-variable */

import { TestBed, async, inject } from '@angular/core/testing';
import { MetricServiceService } from './metric-service.service';

describe('MetricServiceService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [MetricServiceService]
    });
  });

  it('should ...', inject([MetricServiceService], (service: MetricServiceService) => {
    expect(service).toBeTruthy();
  }));
});
