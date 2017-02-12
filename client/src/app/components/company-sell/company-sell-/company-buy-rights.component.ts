import { Component, OnInit } from '@angular/core';
import {TransactionService} from "../../../services/transaction.service";

@Component({
  selector: 'app-company-buy-rights',
  templateUrl: './company-buy-rights.component.html',
  styleUrls: ['./company-buy-rights.component.scss']
})
export class CompanyBuyRightsComponent implements OnInit {
  public leftBuy: string[] = [];
  private transactionService:TransactionService;
  private getAllOnSale:Component;

  constructor (private transactionService:TransactionService){

  }

  ngOnInit() {
    this.leftBuy.push('-26px');
    this.transactionService.getAllonSale().subscribe(onsale => {
      console.log('Get all on sale', onsale);
      this.getAllOnSale = onsale;
    });
  }
}
