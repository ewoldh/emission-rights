import {Component, OnInit} from '@angular/core';
import {CompanyService} from "../../services/company.service";
import {EtaAccountService} from "../../services/eta-account.service";
import {TransactionService} from "../../services/transaction.service";



@Component({
  selector: 'landing',
  templateUrl: 'landing.component.html',
  styleUrls: ['landing.component.scss']
})
export class LandingComponent implements OnInit {
  private companies:Component[];
  private etaAccountUser:Component;
  private buyHistoryById:Component;
  private sellHistoryById:Component;
  private getAllOnSale:Component[];
  private testNumber:number=666;


  constructor(private companyService:CompanyService,private etaAccountService: EtaAccountService, private transactionService: TransactionService ){

  }

  ngOnInit() {
    this.companyService.getAllCompanies().subscribe(companies => {
      console.log('Companies',companies);
      this.companies = companies;
    });

    this.etaAccountService.getEtaAccountUserById().subscribe(etaAccountUser => {
      console.log('etaAccount',etaAccountUser);
      this.etaAccountUser = etaAccountUser;
    });

    this.etaAccountService.postEtaAccount(this.testNumber).subscribe(etaAccountUser => {
      console.log('etaAccount Posting',etaAccountUser);
      this.etaAccountUser = etaAccountUser;
      //console.log('etaAccount Poste',this.etaAccountUser);
      this.etaAccountService.getEtaAccountUserById().subscribe(etaAccountUser => {
        console.log('etaAccount',etaAccountUser);
        this.etaAccountUser = etaAccountUser;
      });
    });

    this.transactionService.getBuyHistoryById().subscribe(transaction => {
      console.log('Bought by id',transaction);
      this.buyHistoryById =transaction;
    });

    this.transactionService.getSellHistoryById().subscribe(transaction => {
      console.log('Sold by id',transaction);
      this.sellHistoryById =transaction;
    });

    this.transactionService.getSellHistoryById().subscribe(onsale => {
      console.log('Get all on sale',onsale);
      this.getAllOnSale =onsale;
    });


  }
}
