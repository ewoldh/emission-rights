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

  @Get('/')
  public getAll(@Req() request: any): any {
    let enrollmentID = new JSONWebToken(request).getUserID();
    console.log('it was here');
    return request.blockchain.query('getAllCompanies', [''], enrollmentID);
  }

}