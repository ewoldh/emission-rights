import { Component, OnInit, Output } from '@angular/core';
import { TransactionService } from "../../../services/transaction.service";


@Component({
  selector: 'app-company-transactions-overview',
  templateUrl: './company-transactions-overview.component.html',
  styleUrls: ['./company-transactions-overview.component.scss']
})
export class CompanyTransactionsOverviewComponent implements OnInit {
@Output() outputData : any[] = [];

  public leftBuy: string[] = [];
  public displayRight: string [] = [];
  private getAllonSale: any;

  constructor(private transactionService: TransactionService) { }

  ngOnInit() {
    this.transactionService.getSellHistoryById().subscribe(onsale => {
      this.getAllonSale = onsale;
      
      this.leftBuy = [];
      for (var item in this.getAllonSale) {
        this.leftBuy.push('-48px');
        this.displayRight.push('0px');
        this.outputData.push([ [false, item] ]);
      };

    });

  }

  ngOnChange() {
    console.log('ngOnChange is activated');
    this.leftBuy = [];
    for (var item in this.getAllonSale) {
      this.leftBuy.push('-48px');
      this.displayRight.push('0px');
      this.outputData.push([ [false, item] ]);
    }
  }

  showDetails(index) {
    this.displayRight[index] = '-380px'
    this.outputData[index] = [true, index];
  }
}