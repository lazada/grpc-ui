import * as $protobuf from "protobufjs";

/** Properties of a ReflectionResponse. */
export interface IReflectionResponse {

    /** ReflectionResponse reflection */
    reflection?: (IReflection|null);

    /** ReflectionResponse error */
    error?: (IError|null);
}

/** Represents a ReflectionResponse. */
export class ReflectionResponse implements IReflectionResponse {

    /**
     * Constructs a new ReflectionResponse.
     * @param [properties] Properties to set
     */
    constructor(properties?: IReflectionResponse);

    /** ReflectionResponse reflection. */
    public reflection?: (IReflection|null);

    /** ReflectionResponse error. */
    public error?: (IError|null);

    /** ReflectionResponse response. */
    public response?: ("reflection"|"error");

    /**
     * Creates a new ReflectionResponse instance using the specified properties.
     * @param [properties] Properties to set
     * @returns ReflectionResponse instance
     */
    public static create(properties?: IReflectionResponse): ReflectionResponse;

    /**
     * Encodes the specified ReflectionResponse message. Does not implicitly {@link ReflectionResponse.verify|verify} messages.
     * @param message ReflectionResponse message or plain object to encode
     * @param [writer] Writer to encode to
     * @returns Writer
     */
    public static encode(message: IReflectionResponse, writer?: $protobuf.Writer): $protobuf.Writer;

    /**
     * Encodes the specified ReflectionResponse message, length delimited. Does not implicitly {@link ReflectionResponse.verify|verify} messages.
     * @param message ReflectionResponse message or plain object to encode
     * @param [writer] Writer to encode to
     * @returns Writer
     */
    public static encodeDelimited(message: IReflectionResponse, writer?: $protobuf.Writer): $protobuf.Writer;

    /**
     * Decodes a ReflectionResponse message from the specified reader or buffer.
     * @param reader Reader or buffer to decode from
     * @param [length] Message length if known beforehand
     * @returns ReflectionResponse
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): ReflectionResponse;

    /**
     * Decodes a ReflectionResponse message from the specified reader or buffer, length delimited.
     * @param reader Reader or buffer to decode from
     * @returns ReflectionResponse
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): ReflectionResponse;

    /**
     * Verifies a ReflectionResponse message.
     * @param message Plain object to verify
     * @returns `null` if valid, otherwise the reason why it is not
     */
    public static verify(message: { [k: string]: any }): (string|null);

    /**
     * Creates a ReflectionResponse message from a plain object. Also converts values to their respective internal types.
     * @param object Plain object
     * @returns ReflectionResponse
     */
    public static fromObject(object: { [k: string]: any }): ReflectionResponse;

    /**
     * Creates a plain object from a ReflectionResponse message. Also converts values to other types if specified.
     * @param message ReflectionResponse
     * @param [options] Conversion options
     * @returns Plain object
     */
    public static toObject(message: ReflectionResponse, options?: $protobuf.IConversionOptions): { [k: string]: any };

    /**
     * Converts this ReflectionResponse to JSON.
     * @returns JSON object
     */
    public toJSON(): { [k: string]: any };
}

/** Properties of a Reflection. */
export interface IReflection {

    /** Reflection service */
    service?: (string[]|null);

    /** Reflection fileDescriptor */
    fileDescriptor?: (Uint8Array[]|null);
}

/** Represents a Reflection. */
export class Reflection implements IReflection {

    /**
     * Constructs a new Reflection.
     * @param [properties] Properties to set
     */
    constructor(properties?: IReflection);

    /** Reflection service. */
    public service: string[];

    /** Reflection fileDescriptor. */
    public fileDescriptor: Uint8Array[];

    /**
     * Creates a new Reflection instance using the specified properties.
     * @param [properties] Properties to set
     * @returns Reflection instance
     */
    public static create(properties?: IReflection): Reflection;

    /**
     * Encodes the specified Reflection message. Does not implicitly {@link Reflection.verify|verify} messages.
     * @param message Reflection message or plain object to encode
     * @param [writer] Writer to encode to
     * @returns Writer
     */
    public static encode(message: IReflection, writer?: $protobuf.Writer): $protobuf.Writer;

    /**
     * Encodes the specified Reflection message, length delimited. Does not implicitly {@link Reflection.verify|verify} messages.
     * @param message Reflection message or plain object to encode
     * @param [writer] Writer to encode to
     * @returns Writer
     */
    public static encodeDelimited(message: IReflection, writer?: $protobuf.Writer): $protobuf.Writer;

    /**
     * Decodes a Reflection message from the specified reader or buffer.
     * @param reader Reader or buffer to decode from
     * @param [length] Message length if known beforehand
     * @returns Reflection
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): Reflection;

    /**
     * Decodes a Reflection message from the specified reader or buffer, length delimited.
     * @param reader Reader or buffer to decode from
     * @returns Reflection
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): Reflection;

    /**
     * Verifies a Reflection message.
     * @param message Plain object to verify
     * @returns `null` if valid, otherwise the reason why it is not
     */
    public static verify(message: { [k: string]: any }): (string|null);

    /**
     * Creates a Reflection message from a plain object. Also converts values to their respective internal types.
     * @param object Plain object
     * @returns Reflection
     */
    public static fromObject(object: { [k: string]: any }): Reflection;

    /**
     * Creates a plain object from a Reflection message. Also converts values to other types if specified.
     * @param message Reflection
     * @param [options] Conversion options
     * @returns Plain object
     */
    public static toObject(message: Reflection, options?: $protobuf.IConversionOptions): { [k: string]: any };

    /**
     * Converts this Reflection to JSON.
     * @returns JSON object
     */
    public toJSON(): { [k: string]: any };
}

/** Properties of an Error. */
export interface IError {

    /** Error message */
    message?: (string|null);
}

/** Represents an Error. */
export class Error implements IError {

    /**
     * Constructs a new Error.
     * @param [properties] Properties to set
     */
    constructor(properties?: IError);

    /** Error message. */
    public message: string;

    /**
     * Creates a new Error instance using the specified properties.
     * @param [properties] Properties to set
     * @returns Error instance
     */
    public static create(properties?: IError): Error;

    /**
     * Encodes the specified Error message. Does not implicitly {@link Error.verify|verify} messages.
     * @param message Error message or plain object to encode
     * @param [writer] Writer to encode to
     * @returns Writer
     */
    public static encode(message: IError, writer?: $protobuf.Writer): $protobuf.Writer;

    /**
     * Encodes the specified Error message, length delimited. Does not implicitly {@link Error.verify|verify} messages.
     * @param message Error message or plain object to encode
     * @param [writer] Writer to encode to
     * @returns Writer
     */
    public static encodeDelimited(message: IError, writer?: $protobuf.Writer): $protobuf.Writer;

    /**
     * Decodes an Error message from the specified reader or buffer.
     * @param reader Reader or buffer to decode from
     * @param [length] Message length if known beforehand
     * @returns Error
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): Error;

    /**
     * Decodes an Error message from the specified reader or buffer, length delimited.
     * @param reader Reader or buffer to decode from
     * @returns Error
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): Error;

    /**
     * Verifies an Error message.
     * @param message Plain object to verify
     * @returns `null` if valid, otherwise the reason why it is not
     */
    public static verify(message: { [k: string]: any }): (string|null);

    /**
     * Creates an Error message from a plain object. Also converts values to their respective internal types.
     * @param object Plain object
     * @returns Error
     */
    public static fromObject(object: { [k: string]: any }): Error;

    /**
     * Creates a plain object from an Error message. Also converts values to other types if specified.
     * @param message Error
     * @param [options] Conversion options
     * @returns Plain object
     */
    public static toObject(message: Error, options?: $protobuf.IConversionOptions): { [k: string]: any };

    /**
     * Converts this Error to JSON.
     * @returns JSON object
     */
    public toJSON(): { [k: string]: any };
}

/** Properties of an InvokeRequest. */
export interface IInvokeRequest {

    /** InvokeRequest method */
    method?: (string|null);

    /** InvokeRequest payload */
    payload?: (Uint8Array|null);
}

/** Represents an InvokeRequest. */
export class InvokeRequest implements IInvokeRequest {

    /**
     * Constructs a new InvokeRequest.
     * @param [properties] Properties to set
     */
    constructor(properties?: IInvokeRequest);

    /** InvokeRequest method. */
    public method: string;

    /** InvokeRequest payload. */
    public payload: Uint8Array;

    /**
     * Creates a new InvokeRequest instance using the specified properties.
     * @param [properties] Properties to set
     * @returns InvokeRequest instance
     */
    public static create(properties?: IInvokeRequest): InvokeRequest;

    /**
     * Encodes the specified InvokeRequest message. Does not implicitly {@link InvokeRequest.verify|verify} messages.
     * @param message InvokeRequest message or plain object to encode
     * @param [writer] Writer to encode to
     * @returns Writer
     */
    public static encode(message: IInvokeRequest, writer?: $protobuf.Writer): $protobuf.Writer;

    /**
     * Encodes the specified InvokeRequest message, length delimited. Does not implicitly {@link InvokeRequest.verify|verify} messages.
     * @param message InvokeRequest message or plain object to encode
     * @param [writer] Writer to encode to
     * @returns Writer
     */
    public static encodeDelimited(message: IInvokeRequest, writer?: $protobuf.Writer): $protobuf.Writer;

    /**
     * Decodes an InvokeRequest message from the specified reader or buffer.
     * @param reader Reader or buffer to decode from
     * @param [length] Message length if known beforehand
     * @returns InvokeRequest
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): InvokeRequest;

    /**
     * Decodes an InvokeRequest message from the specified reader or buffer, length delimited.
     * @param reader Reader or buffer to decode from
     * @returns InvokeRequest
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): InvokeRequest;

    /**
     * Verifies an InvokeRequest message.
     * @param message Plain object to verify
     * @returns `null` if valid, otherwise the reason why it is not
     */
    public static verify(message: { [k: string]: any }): (string|null);

    /**
     * Creates an InvokeRequest message from a plain object. Also converts values to their respective internal types.
     * @param object Plain object
     * @returns InvokeRequest
     */
    public static fromObject(object: { [k: string]: any }): InvokeRequest;

    /**
     * Creates a plain object from an InvokeRequest message. Also converts values to other types if specified.
     * @param message InvokeRequest
     * @param [options] Conversion options
     * @returns Plain object
     */
    public static toObject(message: InvokeRequest, options?: $protobuf.IConversionOptions): { [k: string]: any };

    /**
     * Converts this InvokeRequest to JSON.
     * @returns JSON object
     */
    public toJSON(): { [k: string]: any };
}

/** Properties of an InvokeResponse. */
export interface IInvokeResponse {

    /** InvokeResponse payload */
    payload?: (Uint8Array|null);

    /** InvokeResponse error */
    error?: (IError|null);
}

/** Represents an InvokeResponse. */
export class InvokeResponse implements IInvokeResponse {

    /**
     * Constructs a new InvokeResponse.
     * @param [properties] Properties to set
     */
    constructor(properties?: IInvokeResponse);

    /** InvokeResponse payload. */
    public payload: Uint8Array;

    /** InvokeResponse error. */
    public error?: (IError|null);

    /** InvokeResponse response. */
    public response?: ("payload"|"error");

    /**
     * Creates a new InvokeResponse instance using the specified properties.
     * @param [properties] Properties to set
     * @returns InvokeResponse instance
     */
    public static create(properties?: IInvokeResponse): InvokeResponse;

    /**
     * Encodes the specified InvokeResponse message. Does not implicitly {@link InvokeResponse.verify|verify} messages.
     * @param message InvokeResponse message or plain object to encode
     * @param [writer] Writer to encode to
     * @returns Writer
     */
    public static encode(message: IInvokeResponse, writer?: $protobuf.Writer): $protobuf.Writer;

    /**
     * Encodes the specified InvokeResponse message, length delimited. Does not implicitly {@link InvokeResponse.verify|verify} messages.
     * @param message InvokeResponse message or plain object to encode
     * @param [writer] Writer to encode to
     * @returns Writer
     */
    public static encodeDelimited(message: IInvokeResponse, writer?: $protobuf.Writer): $protobuf.Writer;

    /**
     * Decodes an InvokeResponse message from the specified reader or buffer.
     * @param reader Reader or buffer to decode from
     * @param [length] Message length if known beforehand
     * @returns InvokeResponse
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): InvokeResponse;

    /**
     * Decodes an InvokeResponse message from the specified reader or buffer, length delimited.
     * @param reader Reader or buffer to decode from
     * @returns InvokeResponse
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): InvokeResponse;

    /**
     * Verifies an InvokeResponse message.
     * @param message Plain object to verify
     * @returns `null` if valid, otherwise the reason why it is not
     */
    public static verify(message: { [k: string]: any }): (string|null);

    /**
     * Creates an InvokeResponse message from a plain object. Also converts values to their respective internal types.
     * @param object Plain object
     * @returns InvokeResponse
     */
    public static fromObject(object: { [k: string]: any }): InvokeResponse;

    /**
     * Creates a plain object from an InvokeResponse message. Also converts values to other types if specified.
     * @param message InvokeResponse
     * @param [options] Conversion options
     * @returns Plain object
     */
    public static toObject(message: InvokeResponse, options?: $protobuf.IConversionOptions): { [k: string]: any };

    /**
     * Converts this InvokeResponse to JSON.
     * @returns JSON object
     */
    public toJSON(): { [k: string]: any };
}
