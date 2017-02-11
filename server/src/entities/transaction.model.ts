'use strict';

import * as shortID from 'shortid';

export class Transaction {
  private _transactionID: string;
  private _transactionDate: number;

  public constructor(private _price: number,
                     private _volume: number,
                     private _seller: string,
                     private _buyer: string,
                     private _requestStatus: string,
                     private _transparent: boolean) {
    this._transactionID   = shortID.generate();
    this._transactionDate = new Date().getTime();
  }

  public get transactionID(): string {
    return this._transactionID;
  }

  public get transactionDate(): number {
    return this._transactionDate;
  }

  public get price(): number {
    return this._price;
  }

  public get volume(): number {
    return this._volume;
  }

  public get seller(): string {
    return this._seller;
  }

  public get buyer(): string {
    return this._buyer;
  }

  public get requestStatus(): string {
    return this._requestStatus;
  }

  public get transparent(): boolean {
    return this._transparent;
  }

  public toJSON(): any {
    return {
      'transactionID':   this.transactionID,
      'price':           this.transactionDate,
      'volume':          this.price,
      'transactionDate': this.volume,
      'seller':          this.seller,
      'buyer':           this.buyer,
      'requestStatus':   this.requestStatus,
      'transparent':     this.transparent
    };
  }
}
