export interface Printer {
  tok(text: string): void,
  space(): void,
  newLine(offset: number): void,
}

export class TextPrinter implements Printer {
  buffer: string = '';

  tok(text: string) {
    this.buffer += text;
  }
  space() {
    this.buffer += ' ';
  }

  newLine(offset: number): void {
    this.buffer += '\n';

    for (let i = 0; i < offset; i++) {
      this.buffer += ' ';
    }
  }
}

export function printMessage(message: protobuf.Type, printer: Printer) {
  printer.newLine(0);

  printer.tok('message');
  printer.space();
  printer.tok(message.name);
  printer.space();
  printer.tok('{');

  for (const field of message.fieldsArray) {
    printer.newLine(1);
    printer.tok(field.type);
    printer.space();
    printer.tok(field.name);
    printer.space();
    printer.tok('=');
    printer.space();
    printer.tok('' + field.id);
    printer.tok(';');
  }

  printer.tok('}');
}

