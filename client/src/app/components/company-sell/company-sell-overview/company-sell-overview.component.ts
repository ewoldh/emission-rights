import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-company-sell-overview',
  templateUrl: './company-sell-overview.component.html',
  styleUrls: ['./company-sell-overview.component.scss']
})
export class CompanySellOverviewComponent implements OnInit {
  public leftBuy: string[] = [];

  ngOnInit() {
    this.leftBuy.push('-48px');
  }
}
