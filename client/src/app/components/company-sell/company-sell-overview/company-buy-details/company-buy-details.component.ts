import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-company-buy-details',
  templateUrl: './company-buy-details.component.html',
  styleUrls: ['./company-buy-details.component.scss']
})
export class CompanyBuyDetailsComponent implements OnInit {

  constructor() { }

  ngOnInit() {
  }

  buttonClicked(button: string) {
    switch (button) {
      case 'cancel':
        console.log('clicked cancel');
      case 'buy':
        console.log('clicked buy');
    }
  }

}
