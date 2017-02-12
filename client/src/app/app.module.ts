import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { HttpModule } from '@angular/http';
import { routing, appRoutingProviders } from './app.routing';
import { Configuration } from './app.constants'
import { AuthGuard } from './guards/index';
import { AppComponent } from './app.component';
import { LoginComponent } from './components/login/login.component';
import { ThingsComponent } from './components/things/things.component';
import { ThingService } from './services/thing.service'
import { AuthenticationService } from './services/authentication.service';
import { HeaderComponent } from './components/emission-header/emission-header.component';

import { CompanyBuyComponent } from './components/company-buy/company-buy.component';
import { CompanyBuyFilterComponent } from './components/company-buy/company-buy-filter/company-buy-filter.component';
import { CompanyBuyRightsComponent } from './components/company-buy/company-buy-rights/company-buy-rights.component';
import { CompanyBuyDetailsComponent } from './components/company-buy/company-buy-rights/company-buy-details/company-buy-details.component'

import { CompanySellComponent } from './components/company-sell/company-sell.component';
import { CompanySellNewComponent } from './components/company-sell/company-sell-new/company-sell-new.component';
import { CompanySellOverviewComponent } from './components/company-sell/company-sell-overview/company-sell-overview.component';

import { CompanyTransactionsComponent } from './components/company-transactions/company-transactions.component';
import { CompanyTransactionsOverviewComponent } from './components/company-transactions/company-transactions-overview/company-transactions-overview.component';
import { CompanyTransactionsFilterComponent } from './components/company-transactions/company-transactions-filter/company-transactions-filter.component';

import { LandingComponent } from "./components/landing/landing.component";
import { CompanyService } from "./services/company.service";
import { EtaAccountService } from "./services/eta-account.service";
import { TransactionService } from "./services/transaction.service";

@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    LandingComponent,
    ThingsComponent,
    HeaderComponent,
    CompanyBuyComponent,
    CompanyBuyFilterComponent,
    CompanyBuyRightsComponent,
    CompanyBuyDetailsComponent,
    CompanySellComponent,
    CompanySellNewComponent,
    CompanySellOverviewComponent,
    CompanyTransactionsComponent,
    CompanyTransactionsOverviewComponent,
    CompanyTransactionsFilterComponent
  ],
  imports: [
    BrowserModule,
    FormsModule,
    HttpModule,
    routing
  ],
  providers: [
    appRoutingProviders,
    Configuration,
    AuthenticationService,
    AuthGuard,
    ThingService,
    CompanyService,
    EtaAccountService,
    TransactionService,
  ],
  bootstrap: [AppComponent]
})
export class AppModule {
}
