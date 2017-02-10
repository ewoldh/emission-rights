'use strict';

import * as shortID from 'shortid';

export class ETAAccount {
  private _etaAccountID: string;

  public constructor(private _userID: string,
                     private _companyID: string,
                     private _balance: string,
                     private _amountOfTransactions: string) {
    this._etaAccountID = shortID.generate();
  }

  public get etaAccountID(): string {
    return this._etaAccountID;
  }

  public get userID(): string {
    return this._userID;
  }

  public get companyID(): string {
    return this._companyID;
  }

  public get balance(): string {
    return this._balance;
  }

  public get amountOfTransactions(): string {
    return this._amountOfTransactions;
  }

  public toJSON(): any {
    return {
      'etaAccountID':         this.etaAccountID,
      'userID':               this.userID,
      'companyID':            this.companyID,
      'balance':              this.balance,
      'amountOfTransactions': this.amountOfTransactions
    };
  }
}