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
        this.headerTitle = 'SELLING';
        break;
      case '/landing/buy':
        console.log('yaay');
        this.headerStyles.buying = 'menu-item active';
        this.headerTitle = 'BUYING';
        break;
    }
  }

  changePage(input: string) {
    for (var item in this.headerStyles) {
      this.headerStyles[item] = 'menu-item'
      if (item === input) {
        this.headerStyles[item] ='menu-item active'
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
    }
  }

}
