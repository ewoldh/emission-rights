import {Get, Post, JsonController, Body, Req, UseBefore, Param} from 'routing-controllers';
import {JSONWebToken} from '../../utils/JSONWebToken';
import {UserAuthenticatorMiddleware} from '../../middleware/UserAuthenticatorMiddleware';
import {CORSMiddleware} from '../../middleware/CORSMiddleware';
import {LoggerFactory} from '../../utils/LoggerFactory';
import {Service} from 'typedi';
import {Transaction} from '../../entities/transaction.model';

@JsonController('/transactions')
@UseBefore(UserAuthenticatorMiddleware, CORSMiddleware)
@Service()
export class TransactionController {
  public constructor(private loggerFactory: LoggerFactory) { }

  @Get('/history/sold/')
  public getAllSold(@Req() request: any): any {
    let enrollmentID = new JSONWebToken(request).getUserID();
    return request.blockchain.query('getAllSoldTransactionsByUserID', [''], enrollmentID);
  }

  @Get('/history/bought/')
  public getAllBought(@Req() request: any): any {
    let enrollmentID = new JSONWebToken(request).getUserID();
    return request.blockchain.query('getAllBoughtTransactionsByUserID', [''], enrollmentID);
  }

  @Get('/allOnSale/')
  public getAllOnSale(@Req() request: any): any {
    let enrollmentID = new JSONWebToken(request).getUserID();
    return request.blockchain.query('getAllTransactionsOnSale', [''], enrollmentID);
  }

  @Post('/')
  public post(@Body() transactionTest: any, @Req() request: any): any {
    let transaction = new Transaction(transactionTest.price,
      transactionTest.volume,
      transactionTest.seller,
      transactionTest.dateOfTransaction,
      transactionTest.buyer,
      transactionTest.market,
      transactionTest.requestStatus,
      transactionTest.transparent
    );
  }

  @Post('/finaliseTransaction/:transactionid')
  public postTransaction(@Param('transactionid') TransactionID: string, @Req() request: any): any {
    let enrollmentID = new JSONWebToken(request).getUserID();
    return request.blockchain.invoke('finaliseTransaction', [TransactionID], enrollmentID);
  }
}