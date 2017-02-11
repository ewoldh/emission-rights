'use strict';

import * as crypto from 'crypto';
import {Password} from '../utils/Password';

export class User {
  private _userID: string;
  private _salt: string;
  private _hash: string;
  private _companyID: string;
  private _etaAccountID: string;

  public constructor(userID: string, password: string, companyID: string, etaAccountID: string) {
    this._userID  = userID;
    this._companyID = companyID;
    this._etaAccountID = etaAccountID;
    this._salt    = crypto.randomBytes(16).toString('hex');
    this._hash    = new Password(password, this.salt).toHash();
  }

  public get userID(): string {
    return this._userID;
  }

  public get salt(): string {
    return this._salt;
  }

  public get hash(): string {
    return this._hash;
  }

  public get companyID(): string {
    return this._companyID;
  }

  public get etaAccountID(): string {
    return this._etaAccountID;
  }

  public toJSON(): any {
    return {
      'userID':  this.userID,
      'salt':    this.salt,
      'hash':    this.hash,
      'companyID': this.companyID,
      'etaAccountID': this.etaAccountID
    };
  }
}