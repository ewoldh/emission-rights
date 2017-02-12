import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'emission-header',
  templateUrl: 'emission-header.component.html',
  styleUrls: ['emission-header.component.scss']
})
export class HeaderComponent implements OnInit {
  public headerStyles = {
    profile: 'menu-item',
    selling: 'menu-item',
    buying: 'menu-item',
    transactions: 'menu-item'
  };

  public headerTitle: string = '';

  constructor(private router: Router ) { }

  ngOnInit() {
    switch(this.router.url) {
      case '/landing/sell':
        this.headerStyles.selling = 'menu-item active';
        this.headerStyles.selling = 'menu-item active';
        break;
      case '/landing/buy':
        this.headerStyles.buying = 'menu-item active';
        this.headerTitle = 'BUYING';
        break;
      case '/landing/transactions':
        this.headerTitle = 'TRANSACTIONS';
        this.headerStyles.transactions = 'menu-item active';
        break;
      case '/landing/profile':
        this.headerTitle = 'PROFILE';
        this.headerStyles.profile = 'menu-item active';
        break;
    }
  }

  changePage(input: string) {
    for (var item in this.headerStyles) {
      if (item === input) {
        this.headerStyles[item] ='menu-item active';
      } else {
        this.headerStyles[item] = 'menu-item';
      }
    }

    switch(input) {
      case 'selling':
        this.headerTitle = 'SELLING';
        this.router.navigate(['./landing/sell']);
        break;
      case 'buying':
        this.headerTitle = 'BUYING';
        this.router.navigate(['./landing/buy']);
        break;
      case 'transactions':
        this.headerTitle = 'TRANSACTIONS';
        this.router.navigate(['./landing/transactions']);
        break;
      case 'profile':
        this.headerTitle = 'PROFILE';
        this.router.navigate(['./landing/profile']);
        break;

    }
  }

}
