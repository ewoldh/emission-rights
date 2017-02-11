import {Component, OnInit} from '@angular/core';

@Component({
  selector: 'landing',
  templateUrl: 'landing.component.html',
  styleUrls: ['landing.component.scss']
})
export class LandingComponent implements OnInit {

  ngOnInit() {
    console.log('LandingComponent Initialized');
  }

}
