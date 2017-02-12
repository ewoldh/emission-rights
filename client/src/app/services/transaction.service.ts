import {Injectable} from '@angular/core';
import {Http} from '@angular/http';
import {Configuration} from '../app.constants';
import {AuthenticationService} from './authentication.service';

@Injectable()
export class TransactionService {

  private actionUrlBought: string;
  private actionUrlSold: string;
  private actionUrlAllOnSale: string;
  private actionUrlFinalise: string;
  private actionUrl: string;
  private headers: any;

  constructor(private _http: Http,
              private _configuration: Configuration,
              private _authenticationService: AuthenticationService) {
    this.actionUrlBought    = _configuration.Server + 'api/v1/transactions/history/bought';
    this.actionUrlSold      = _configuration.Server + 'api/v1/transactions/history/sold';
    this.actionUrlAllOnSale = _configuration.Server + 'api/v1/transactions/allOnSale';
    this.actionUrl          = _configuration.Server + 'api/v1/transactions';
    this.actionUrlFinalise  = _configuration.Server + 'api/v1/transactions/finaliseTrade';
    this.headers            = _authenticationService.createAuthorizationHeader();
  }

  public getBuyHistoryById() {
    let user: any = JSON.parse(localStorage.getItem('currentUser')).user;
    return this._http
      .get(this.actionUrlBought, {headers: this.headers})
      .map(res => res.json());
  }

  public getSellHistoryById() {
    let user: any = JSON.parse(localStorage.getItem('currentUser')).user;
    return this._http
      .get(this.actionUrlSold, {headers: this.headers})
      .map(res => res.json());
  }

  public getAllonSale() {
    let user: any = JSON.parse(localStorage.getItem('currentUser')).user;
    return this._http
      .get(this.actionUrlAllOnSale, {headers: this.headers})
      .map(res => res.json());
  }

  public postTransaction(transaction: any) {
    return this._http.post(this.actionUrl, transaction, {headers: this.headers})
      .map(res => res.json());
  }

  public postFinaliseTransaction(transactionID: number) {
    return this._http.post(this.actionUrlFinalise, {transactionID}, {headers: this.headers})
      .map(res => res.json());
  }
}
