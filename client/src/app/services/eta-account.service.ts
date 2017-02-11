import {Injectable} from '@angular/core';
import {Http} from '@angular/http';
import {Configuration} from '../app.constants';
import {AuthenticationService} from './authentication.service';

@Injectable()
export class EtaAccountService {
  private actionUrl: string;
  private headers: any;

  constructor(private _http: Http,
              private _configuration: Configuration,
              private _authenticationService: AuthenticationService) {
    this.actionUrl = _configuration.Server + 'api/v1/etaAccounts';
    this.headers   = _authenticationService.createAuthorizationHeader();
  }

  public getEtaAccountUserById() {
    let user: any = JSON.parse(localStorage.getItem('currentUser')).user;
    return this._http
      .get(this.actionUrl + '/' + user.userID, {headers: this.headers})
      .map(res => res.json());
  }

  public postEtaAccount(amountValue: number) {
    return this._http.post(this.actionUrl, {amount: amountValue}, {headers: this.headers})
      .map(res => res.json());
  }

}
