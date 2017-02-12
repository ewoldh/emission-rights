import { ModuleWithProviders } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { AuthGuard } from './guards/index';
import { LoginComponent } from './components/login/login.component';
import { ThingsComponent } from './components/things/things.component';
import { CompanyBuyComponent } from './components/company-buy/company-buy.component';
import { CompanySellComponent} from './components/company-sell/company-sell.component';
import { CompanyTransactionsComponent} from './components/company-transactions/company-transactions.component';
import { LandingComponent } from "./components/landing/landing.component";

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
      path: 'sell',
      component: CompanySellComponent
    },
    {
      path: 'transactions',
      component: CompanyTransactionsComponent
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
