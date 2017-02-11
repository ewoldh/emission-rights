'use strict';

import * as crypto from 'crypto';
import {Password} from '../utils/Password';

export class User {
  private _userID: string;
  private _salt: string;
  private _hash: string;
  private _company: string;

  public constructor(userID: string, password: string, company: string) {
    this._userID  = userID;
    this._company = company;
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

  public get company(): string {
    return this._company;
  }

  public toJSON(): any {
    return {
      'userID':  this.userID,
      'salt':    this.salt,
      'hash':    this.hash,
      'company': this.company
    };
  }
}