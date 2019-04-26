/*
 * ORY Hydra
 * Welcome to the ORY Hydra HTTP API documentation. You will find documentation for all HTTP APIs here.
 *
 * OpenAPI spec version: latest
 * 
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 * Do not edit the class manually.
 */


package com.github.ory.hydra.model;

import java.util.Objects;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonValue;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

/**
 * LogoutRequest
 */
@javax.annotation.Generated(value = "io.swagger.codegen.languages.JavaClientCodegen", date = "2019-04-26T09:49:08.655+02:00")
public class LogoutRequest {
  @JsonProperty("request_url")
  private String requestUrl = null;

  @JsonProperty("rp_initiated")
  private Boolean rpInitiated = null;

  @JsonProperty("sid")
  private String sid = null;

  @JsonProperty("subject")
  private String subject = null;

  public LogoutRequest requestUrl(String requestUrl) {
    this.requestUrl = requestUrl;
    return this;
  }

   /**
   * RequestURL is the original Logout URL requested.
   * @return requestUrl
  **/
  @ApiModelProperty(value = "RequestURL is the original Logout URL requested.")
  public String getRequestUrl() {
    return requestUrl;
  }

  public void setRequestUrl(String requestUrl) {
    this.requestUrl = requestUrl;
  }

  public LogoutRequest rpInitiated(Boolean rpInitiated) {
    this.rpInitiated = rpInitiated;
    return this;
  }

   /**
   * RPInitiated is set to true if the request was initiated by a Relying Party (RP), also known as an OAuth 2.0 Client.
   * @return rpInitiated
  **/
  @ApiModelProperty(value = "RPInitiated is set to true if the request was initiated by a Relying Party (RP), also known as an OAuth 2.0 Client.")
  public Boolean getRpInitiated() {
    return rpInitiated;
  }

  public void setRpInitiated(Boolean rpInitiated) {
    this.rpInitiated = rpInitiated;
  }

  public LogoutRequest sid(String sid) {
    this.sid = sid;
    return this;
  }

   /**
   * SessionID is the login session ID that was requested to log out.
   * @return sid
  **/
  @ApiModelProperty(value = "SessionID is the login session ID that was requested to log out.")
  public String getSid() {
    return sid;
  }

  public void setSid(String sid) {
    this.sid = sid;
  }

  public LogoutRequest subject(String subject) {
    this.subject = subject;
    return this;
  }

   /**
   * Subject is the user for whom the logout was request.
   * @return subject
  **/
  @ApiModelProperty(value = "Subject is the user for whom the logout was request.")
  public String getSubject() {
    return subject;
  }

  public void setSubject(String subject) {
    this.subject = subject;
  }


  @Override
  public boolean equals(java.lang.Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    LogoutRequest logoutRequest = (LogoutRequest) o;
    return Objects.equals(this.requestUrl, logoutRequest.requestUrl) &&
        Objects.equals(this.rpInitiated, logoutRequest.rpInitiated) &&
        Objects.equals(this.sid, logoutRequest.sid) &&
        Objects.equals(this.subject, logoutRequest.subject);
  }

  @Override
  public int hashCode() {
    return Objects.hash(requestUrl, rpInitiated, sid, subject);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class LogoutRequest {\n");
    
    sb.append("    requestUrl: ").append(toIndentedString(requestUrl)).append("\n");
    sb.append("    rpInitiated: ").append(toIndentedString(rpInitiated)).append("\n");
    sb.append("    sid: ").append(toIndentedString(sid)).append("\n");
    sb.append("    subject: ").append(toIndentedString(subject)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(java.lang.Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }
  
}

