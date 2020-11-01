/* eslint-disable */

import { UnaryInterceptor, Request } from "grpc-web"

/**
 * @object
 * @implements {UnaryInterceptor}
 */
export class AuthUnaryInterceptor implements UnaryInterceptor<any, any> {
    token?: string | undefined
  
    constructor(token: string | undefined) {
      this.token = token
    }

    intercept(request: Request<any, any>, invoker: any) {
        // Update the request message before the RPC.
        const metadata = request.getMetadata()
        metadata.Authorization = 'Bearer ' + this.token
        console.log('Added bearer '+ this.token)
        return invoker(request)
    }
}