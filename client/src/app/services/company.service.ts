import { Injectable } from '@angular/core';
import {Http} from "@angular/http";
import {Configuration} from "../app.constants";
import {AuthenticationService} from "./authentication.service";
import 'rxjs/add/operator/map';

@Injectable()
export class CompanyService {
  private actionUrl: string;
  private headers: any;

  constructor(private _http: Http,
              private _configuration: Configuration,
              private _authenticationService: AuthenticationService) {
    this.actionUrl = _configuration.Server + 'api/v1/etaAccounts';
    this.headers = _authenticationService.createAuthorizationHeader();
  }

  public getAllCompanies() {
    return this._http
      .get(this.actionUrl, {headers: this.headers})
      .map(res => res.json());
  }
}
