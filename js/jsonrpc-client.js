export class JSONRPCError extends Error {
  constructor(message, name, code, data) {
    super(message);
    this.name = name;
    this.code = code;
    this.data = data;
  }
}

class JSONRPCScheduler {
  /**
   *
   * @param {*} transport
   */
  constructor(transport) {
    this._transport = transport;
    this._requestID = 0;
    this._scheduleRequests = {};
    this._commitTimerID = null;
    this._beforeRequest = null;
  }
  beforeRequest(fn) {
    this._beforeRequest = fn;
  }
  __scheduleCommit() {
    if (this._commitTimerID) {
      clearTimeout(this._commitTimerID);
    }
    this._commitTimerID = setTimeout(() => {
      this._commitTimerID = null;
      const scheduleRequests = { ...this._scheduleRequests };
      this._scheduleRequests = {};
      let requests = [];
      for (let key in scheduleRequests) {
        requests.push(scheduleRequests[key].request);
      }
      this.__doRequest(requests)
        .then((responses) => {
          for (let i = 0; i < responses.length; i++) {
            const schedule = scheduleRequests[responses[i].id];
            if (responses[i].error) {
              schedule.reject(responses[i].error);
              continue;
            }
            schedule.resolve(responses[i].result);
          }
        })
        .catch((e) => {
          for (let key in requests) {
            if (!requests.hasOwnProperty(key)) {
              continue;
            }
            if (scheduleRequests.hasOwnProperty(requests[key].id)) {
              scheduleRequests[requests[key].id].reject(e);
            }
          }
        });
    }, 0);
  }
  makeJSONRPCRequest(id, method, params) {
    return {
      jsonrpc: "2.0",
      id: id,
      method: method,
      params: params,
    };
  }
  /**
   * @param {string} method
   * @param {Object} params
   * @returns {Promise<*>}
   */
  __scheduleRequest(method, params) {
    const p = new Promise((resolve, reject) => {
      const request = this.makeJSONRPCRequest(
        this.__requestIDGenerate(),
        method,
        params
      );
      this._scheduleRequests[request.id] = {
        request,
        resolve,
        reject,
      };
    });
    this.__scheduleCommit();
    return p;
  }
  __doRequest(request) {
    return this._transport.doRequest(request);
  }
  __requestIDGenerate() {
    return ++this._requestID;
  }
}
class JSONRPCClientUser {
  constructor(transport) {
    this.scheduler = new JSONRPCScheduler(transport);
  }

  /**
   * @param {number} id
   * @return {PromiseLike<{Name: string}>}
   **/
  getUserNameByID(id) {
    return this.scheduler
      .__scheduleRequest("user.getUserNameByID", { id: id })
      .catch((e) => {
        throw userGetUserNameByIDConvertError(e);
      });
  }
}

class JSONRPCClient {
  constructor(transport) {
    this.user = new JSONRPCClientUser(transport);
  }
}
export default JSONRPCClient;

function userGetUserNameByIDConvertError(e) {
  switch (e.code) {
    default:
      return new JSONRPCError(e.message, "UnknownError", e.code, e.data);
  }
}
