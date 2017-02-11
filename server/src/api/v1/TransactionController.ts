import {Get, Post, JsonController, Param, Body, Req, UseBefore} from 'routing-controllers';
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

  @Get('/history/sold/:id')
  public getAllSold(@Param('id') userID: string, @Req() request: any): any {
    let enrollmentID = new JSONWebToken(request).getUserID();
    return request.blockchain.query('getAllSoldTransactionsByUserID', [userID], enrollmentID);
  }

  @Get('/history/bought/:id')
  public getAllBought(@Param('id') userID: string, @Req() request: any): any {
    let enrollmentID = new JSONWebToken(request).getUserID();
    return request.blockchain.query('getAllBoughtTransactionsByUserID', [userID], enrollmentID);
  }

  @Get('/allOnSale/')
  public getAllOnSale(@Req() request: any): any {
    let enrollmentID = new JSONWebToken(request).getUserID();
    return request.blockchain.query('getAllTransactionsOnSale', [''], enrollmentID);
  }

  @Post('/')
  public post(@Body() transaction: Transaction, @Req() request: any): any {
    let enrollmentID = new JSONWebToken(request).getUserID();
    return request.blockchain.invoke('createTransaction', [transaction.toJSON()], enrollmentID);
  }
}