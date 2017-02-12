import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-company-buy-rights',
  templateUrl: './company-buy-rights.component.html',
  styleUrls: ['./company-buy-rights.component.scss']
})
export class CompanyBuyRightsComponent implements OnInit {
  public leftBuy: string[] = [];

  ngOnInit() {
    this.leftBuy.push('-26px');
  }
}
