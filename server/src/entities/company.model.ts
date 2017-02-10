'use strict';

import * as shortID from 'shortid';

export class Company {
    private _companyID: string;

    public constructor(private _username: string,
                       private _companyName: string,
                       private _branch: string,
                       private _approvalStatus: string) {
        this._companyID = shortID.generate();
    }

    public get companyID(): string {
        return this._companyID;
    }

    public get username(): string {
        return this._username;
    }

    public get companyName(): string {
        return this._companyName;
    }

    public get branch(): string {
        return this._branch;
    }

    public get approvalStatus(): string {
        return this._approvalStatus;
    }

    public toJSON(): any {
        return {
            'thingID': this.companyID,
            'someProperty': this.username,
            'userID': this.companyName,
            'accountBalance': this.branch,
            'approvalStatus': this.approvalStatus
        };
    }
}