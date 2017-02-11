'use strict';

import * as shortID from 'shortid';

export class BankAccount {
    private _bankAccountID: string;

    public constructor(private _companyID: string,
                       private _userID: string,
                       private _accountBalance: number) {
        this._bankAccountID = shortID.generate();
    }

    public get bankAccountID(): string {
        return this._bankAccountID;
    }

    public get companyID(): string {
        return this._companyID;
    }

    public get userID(): string {
        return this._userID;
    }

    public get accountBalance(): number {
        return this._accountBalance;
    }

    public toJSON(): any {
        return {
            'thingID': this.bankAccountID,
            'someProperty': this.companyID,
            'userID': this.userID,
            'accountBalance': this.accountBalance
        };
    }
}