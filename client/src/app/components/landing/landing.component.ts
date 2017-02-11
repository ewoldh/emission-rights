import {Component, OnInit} from '@angular/core';
import {CompanyService} from "../../services/company.service";
import {EtaAccountService} from "../../services/eta-account.service";



@Component({
  selector: 'landing',
  templateUrl: 'landing.component.html',
  styleUrls: ['landing.component.scss']
})
export class LandingComponent implements OnInit {
  private companies:Component[];
  private etaAccountUser:Component;

  constructor(private companyService:CompanyService,private etaAccountService: EtaAccountService){

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
  }
}
