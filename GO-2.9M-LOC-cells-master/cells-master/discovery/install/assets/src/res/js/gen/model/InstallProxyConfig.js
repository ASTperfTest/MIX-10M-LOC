/**
 * Pydio Cells Rest API
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: 1.0
 * 
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 * Do not edit the class manually.
 *
 */


import ApiClient from '../ApiClient';
import InstallTLSCertificate from './InstallTLSCertificate';
import InstallTLSLetsEncrypt from './InstallTLSLetsEncrypt';
import InstallTLSSelfSigned from './InstallTLSSelfSigned';





/**
* The InstallProxyConfig model module.
* @module model/InstallProxyConfig
* @version 1.0
*/
export default class InstallProxyConfig {
    /**
    * Constructs a new <code>InstallProxyConfig</code>.
    * @alias module:model/InstallProxyConfig
    * @class
    */

    constructor() {
        

        
        

        

        
    }

    /**
    * Constructs a <code>InstallProxyConfig</code> from a plain JavaScript object, optionally creating a new instance.
    * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
    * @param {Object} data The plain JavaScript object bearing properties of interest.
    * @param {module:model/InstallProxyConfig} obj Optional instance to populate.
    * @return {module:model/InstallProxyConfig} The populated <code>InstallProxyConfig</code> instance.
    */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InstallProxyConfig();

            
            
            

            if (data.hasOwnProperty('BindURL')) {
                obj['BindURL'] = ApiClient.convertToType(data['BindURL'], 'String');
            }
            if (data.hasOwnProperty('ExternalURL')) {
                obj['ExternalURL'] = ApiClient.convertToType(data['ExternalURL'], 'String');
            }
            if (data.hasOwnProperty('RedirectURLs')) {
                obj['RedirectURLs'] = ApiClient.convertToType(data['RedirectURLs'], ['String']);
            }
            if (data.hasOwnProperty('SelfSigned')) {
                obj['SelfSigned'] = InstallTLSSelfSigned.constructFromObject(data['SelfSigned']);
            }
            if (data.hasOwnProperty('LetsEncrypt')) {
                obj['LetsEncrypt'] = InstallTLSLetsEncrypt.constructFromObject(data['LetsEncrypt']);
            }
            if (data.hasOwnProperty('Certificate')) {
                obj['Certificate'] = InstallTLSCertificate.constructFromObject(data['Certificate']);
            }
        }
        return obj;
    }

    /**
    * @member {String} BindURL
    */
    BindURL = undefined;
    /**
    * @member {String} ExternalURL
    */
    ExternalURL = undefined;
    /**
    * @member {Array.<String>} RedirectURLs
    */
    RedirectURLs = undefined;
    /**
    * @member {module:model/InstallTLSSelfSigned} SelfSigned
    */
    SelfSigned = undefined;
    /**
    * @member {module:model/InstallTLSLetsEncrypt} LetsEncrypt
    */
    LetsEncrypt = undefined;
    /**
    * @member {module:model/InstallTLSCertificate} Certificate
    */
    Certificate = undefined;








}


