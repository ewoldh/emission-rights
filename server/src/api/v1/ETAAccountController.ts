import {Get, JsonController, Param, Req, UseBefore, Post, Body} from 'routing-controllers';
import {JSONWebToken} from '../../utils/JSONWebToken';
import {UserAuthenticatorMiddleware} from '../../middleware/UserAuthenticatorMiddleware';
import {CORSMiddleware} from '../../middleware/CORSMiddleware';
import {LoggerFactory} from '../../utils/LoggerFactory';
import {Service} from 'typedi';

class Value {
  public amount: number;
}

@JsonController('/etaAccounts')
@UseBefore(UserAuthenticatorMiddleware, CORSMiddleware)
@Service()
export class ETAAccountController {
  public constructor(private loggerFactory: LoggerFactory) { }

  @Get('/:id')
  public getETAAccountByUserID(@Param('id') userID: string, @Req() request: any): any {
    let enrollmentID = new JSONWebToken(request).getUserID();
    return request.blockchain.query('getETAAccountByUserID', [userID], enrollmentID);
  }

  @Post('/')
  public addETAToTradePlatform(@Body() amount: number, @Req() request: any): any {
    let enrollmentID = new JSONWebToken(request).getUserID();
    return request.blockchain.invoke('createETAs', [amount.toString()], enrollmentID);
  }
}