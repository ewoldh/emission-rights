import {Get, JsonController, Req, UseBefore} from 'routing-controllers';
import {JSONWebToken} from '../../utils/JSONWebToken';
import {UserAuthenticatorMiddleware} from '../../middleware/UserAuthenticatorMiddleware';
import {CORSMiddleware} from '../../middleware/CORSMiddleware';
import {LoggerFactory} from '../../utils/LoggerFactory';
import {Service} from 'typedi';

@JsonController('/companies')
@UseBefore(UserAuthenticatorMiddleware, CORSMiddleware)
@Service()
export class CompaniesController {
  public constructor(private loggerFactory: LoggerFactory) { }

  @Get('/all')
  public getAll(@Req() request: any): any {
    let enrollmentID = new JSONWebToken(request).getUserID();
    return request.blockchain.query('getAllCompanies', [''], enrollmentID);
  }

}