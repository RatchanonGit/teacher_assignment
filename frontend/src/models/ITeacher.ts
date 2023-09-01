import { StudentsInterface } from "./IStudent";

export interface TeachersInterface {

    ID?: number,
    Faculty_ID?: number;
    Level?: string;
    Name?: string;
    Email?: string;
    Graduate_faculty_level_ID?: number;
    Officer_ID?: number;

    Student?: StudentsInterface
   }
