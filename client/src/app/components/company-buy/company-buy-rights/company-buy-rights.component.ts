import { Component, OnInit, Output } from '@angular/core';
import {TransactionService} from "../../../services/transaction.service";

@Component({
  selector: 'app-company-buy-rights',
  templateUrl: './company-buy-rights.component.html',
  styleUrls: ['./company-buy-rights.component.scss']
})
export class CompanyBuyRightsComponent implements OnInit {
  
  @Output() outputData : any[] = [];

  public leftBuy: string[] = [];
  private getAllonSale: any;

  constructor(private transactionService: TransactionService) { }

  ngOnInit() {
    this.transactionService.getAllonSale().subscribe(onsale => {
      this.getAllonSale = onsale;
      
      this.leftBuy = [];
      for (var item in this.getAllonSale) {
        this.leftBuy.push('-26px');
        this.outputData.push([ [false, item] ]);
      };

    });

  }

  ngOnChange() {
    console.log('ngOnChange is activated');
    this.leftBuy = [];
    for (var item in this.getAllonSale) {
      this.leftBuy.push('-26px');
      this.outputData.push([ [false, item] ]);
    }
  }

  showDetails(index) {
    this.outputData[index] = [true, index];
  }
}