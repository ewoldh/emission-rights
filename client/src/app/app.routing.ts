import {ModuleWithProviders} from '@angular/core';
import {Routes, RouterModule} from '@angular/router';
import {AuthGuard} from './guards/index';
import {LoginComponent} from './components/login/login.component';
import {ThingsComponent} from './components/things/things.component';
import {CompanyBuyComponent} from './components/company-buy/company-buy.component';
import {LandingComponent} from "./components/landing/landing.component";

const appRoutes: Routes = [
  {path: 'login', component: LoginComponent},
  {path: 'landing', component: LandingComponent, //canActivate: [AuthGuard],
  children: [
    {
      path: '',
      redirectTo: 'buy',
      pathMatch: 'full'
    },
    {
      path: 'buy',
      component: CompanyBuyComponent
    },
    {
      path: 'things',
      component: ThingsComponent
    }]
  },

  // otherwise redirect to login
  {path: '**', redirectTo: 'login'}
];

export const appRoutingProviders: any[] = [];

export const routing: ModuleWithProviders = RouterModule.forRoot(appRoutes);
