/**
 * ORY Hydra
 * Welcome to the ORY Hydra HTTP API documentation. You will find documentation for all HTTP APIs here.
 *
 * OpenAPI spec version: latest
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 *
 * Swagger Codegen version: 2.2.3
 *
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD. Register as an anonymous module.
    define(['ApiClient'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'));
  } else {
    // Browser globals (root is window)
    if (!root.OryHydra) {
      root.OryHydra = {};
    }
    root.OryHydra.OpenIDConnectContext = factory(root.OryHydra.ApiClient);
  }
}(this, function(ApiClient) {
  'use strict';




  /**
   * The OpenIDConnectContext model module.
   * @module model/OpenIDConnectContext
   * @version latest
   */

  /**
   * Constructs a new <code>OpenIDConnectContext</code>.
   * @alias module:model/OpenIDConnectContext
   * @class
   */
  var exports = function() {
    var _this = this;






  };

  /**
   * Constructs a <code>OpenIDConnectContext</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/OpenIDConnectContext} obj Optional instance to populate.
   * @return {module:model/OpenIDConnectContext} The populated <code>OpenIDConnectContext</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();

      if (data.hasOwnProperty('acr_values')) {
        obj['acr_values'] = ApiClient.convertToType(data['acr_values'], ['String']);
      }
      if (data.hasOwnProperty('display')) {
        obj['display'] = ApiClient.convertToType(data['display'], 'String');
      }
      if (data.hasOwnProperty('id_token_hint_claims')) {
        obj['id_token_hint_claims'] = ApiClient.convertToType(data['id_token_hint_claims'], {'String': Object});
      }
      if (data.hasOwnProperty('login_hint')) {
        obj['login_hint'] = ApiClient.convertToType(data['login_hint'], 'String');
      }
      if (data.hasOwnProperty('ui_locales')) {
        obj['ui_locales'] = ApiClient.convertToType(data['ui_locales'], ['String']);
      }
    }
    return obj;
  }

  /**
   * ACRValues is the Authentication AuthorizationContext Class Reference requested in the OAuth 2.0 Authorization request. It is a parameter defined by OpenID Connect and expresses which level of authentication (e.g. 2FA) is required.  OpenID Connect defines it as follows: > Requested Authentication AuthorizationContext Class Reference values. Space-separated string that specifies the acr values that the Authorization Server is being requested to use for processing this Authentication Request, with the values appearing in order of preference. The Authentication AuthorizationContext Class satisfied by the authentication performed is returned as the acr Claim Value, as specified in Section 2. The acr Claim is requested as a Voluntary Claim by this parameter.
   * @member {Array.<String>} acr_values
   */
  exports.prototype['acr_values'] = undefined;
  /**
   * Display is a string value that specifies how the Authorization Server displays the authentication and consent user interface pages to the End-User. The defined values are: page: The Authorization Server SHOULD display the authentication and consent UI consistent with a full User Agent page view. If the display parameter is not specified, this is the default display mode. popup: The Authorization Server SHOULD display the authentication and consent UI consistent with a popup User Agent window. The popup User Agent window should be of an appropriate size for a login-focused dialog and should not obscure the entire window that it is popping up over. touch: The Authorization Server SHOULD display the authentication and consent UI consistent with a device that leverages a touch interface. wap: The Authorization Server SHOULD display the authentication and consent UI consistent with a \"feature phone\" type display.  The Authorization Server MAY also attempt to detect the capabilities of the User Agent and present an appropriate display.
   * @member {String} display
   */
  exports.prototype['display'] = undefined;
  /**
   * IDTokenHintClaims are the claims of the ID Token previously issued by the Authorization Server being passed as a hint about the End-User's current or past authenticated session with the Client.
   * @member {Object.<String, Object>} id_token_hint_claims
   */
  exports.prototype['id_token_hint_claims'] = undefined;
  /**
   * LoginHint hints about the login identifier the End-User might use to log in (if necessary). This hint can be used by an RP if it first asks the End-User for their e-mail address (or other identifier) and then wants to pass that value as a hint to the discovered authorization service. This value MAY also be a phone number in the format specified for the phone_number Claim. The use of this parameter is optional.
   * @member {String} login_hint
   */
  exports.prototype['login_hint'] = undefined;
  /**
   * UILocales is the End-User'id preferred languages and scripts for the user interface, represented as a space-separated list of BCP47 [RFC5646] language tag values, ordered by preference. For instance, the value \"fr-CA fr en\" represents a preference for French as spoken in Canada, then French (without a region designation), followed by English (without a region designation). An error SHOULD NOT result if some or all of the requested locales are not supported by the OpenID Provider.
   * @member {Array.<String>} ui_locales
   */
  exports.prototype['ui_locales'] = undefined;



  return exports;
}));


