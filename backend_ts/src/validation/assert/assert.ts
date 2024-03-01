import { ValidationError } from "../error";
import { Constraints, Type } from "../parser.type";

export const isNot = (value: unknown, type: Type): ValidationError => {
  return new ValidationError(`${value} is not of type ${type}`);
};

export const isNotUnion = <T>(value: unknown, values: T[]): ValidationError => {
  return new ValidationError(`${value} is not one of ${values.join(" | ")}`);
};

export const isNotMin = (value: number, min: number): ValidationError => {
  return new ValidationError(`${value} must not be smaller than ${min}`);
};

export const isNotMax = (value: number, max: number): ValidationError => {
  return new ValidationError(`${value} must not be larger than ${max}`);
};

export const isNotMinLength = (value: string, min: number): ValidationError => {
  return new ValidationError(
    `${value} must be at least ${min} characters long`
  );
};

export const isNotMaxLength = (value: string, max: number): ValidationError => {
  return new ValidationError(
    `${value} must not be longer than ${max} characters`
  );
};

export function isString(
  value: unknown,
  constraints: Constraints
): asserts value is string {
  if (typeof value !== "string") throw isNot(value, "string");
  if (constraints.min && value.length < constraints.min)
    throw isNotMinLength(value, constraints.min);
  if (constraints.max && value.length > constraints.max)
    throw isNotMaxLength(value, constraints.max);
}

export function isNumber(
  value: unknown,
  constraints: Constraints
): asserts value is number {
  if (typeof value !== "number") throw isNot(value, "number");
  if (Number.isNaN(value)) throw isNot(value, "number");
  if (constraints.min && value < constraints.min)
    throw isNotMin(value, constraints.min);
  if (constraints.max && value > constraints.max)
    throw isNotMax(value, constraints.max);
}

export function isObject(value: unknown): asserts value is object {
  if (!value) throw isNot(value, "object");
  if (typeof value !== "object") throw isNot(value, "object");
  if (Array.isArray(value)) throw isNot(value, "object");
}
