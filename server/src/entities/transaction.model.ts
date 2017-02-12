'use strict';

import * as shortID from 'shortid';

export class Transaction {
  private _transactionID: string;
  private _sellDate: number;

  public constructor(private _price: number,
                     private _volume: number,
                     private _seller: string,
                     private _dateOfTransaction: number,
                     private _buyer: string,
                     private _market: string,
                     private _requestStatus: string,
                     private _transparent: boolean) {
    this._transactionID = shortID.generate();
    this._sellDate      = new Date().getTime();
  }

  public get transactionID(): string {
    return this._transactionID;
  }

  public get sellDate(): number {
    return this._sellDate;
  }

  public get price(): number {
    return this._price;
  }

  public get volume(): number {
    return this._volume;
  }

  public get dateOfTransaction(): number {
    return this._dateOfTransaction;
  }

  public get seller(): string {
    return this._seller;
  }

  public get buyer(): string {
    return this._buyer;
  }

  public get market(): string {
    return this._market;
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
      'price':           this.sellDate,
      'volume':          this.price,
      'transactionDate': this.volume,
      'buyDate':         this.dateOfTransaction,
      'seller':          this.seller,
      'buyer':           this.buyer,
      'market':          this.market,
      'requestStatus':   this.requestStatus,
      'transparent':     this.transparent
    };
  }
}
