import { Component, OnInit, Input } from '@angular/core';
import {TransactionService} from "../../../../services/transaction.service";

@Component({
  selector: 'app-company-buy-details',
  templateUrl: './company-buy-details.component.html',
  styleUrls: ['./company-buy-details.component.scss']
})
export class CompanyBuyDetailsComponent implements OnInit {
  @Input() inputData: any [];
  public getSpecificData: any;

  constructor(private transactionService: TransactionService) { }

  ngOnInit() {
    this.inputData[0] = false;
    this.transactionService.getAllonSale().subscribe(onsale => {
        this.getSpecificData = onsale;
        
        for (var item in onsale) {
          this.getSpecificData.push(item);
        }
    });
  }

  buttonClicked(button: string) {
    switch (button) {
      case 'cancel':
        this.inputData[0] = false;
        console.log(this.inputData);
        break;
      case 'buy':
        console.log(this.getSpecificData[this.inputData[1]]);
        break;
    }
  }
}
