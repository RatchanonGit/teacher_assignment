import { StudentsInterface } from "./IStudent";
import { TeachersInterface } from "./ITeacher";
import { Teaching_durationsInterface } from "./ITeaching_duration";
import { Content_difficulty_levelsInterface } from "./IContent_difficulty_level";
import { Content_qualitysInterface } from "./IContent_quality";
import { CommentsInterface } from "./IComment";

export interface Teacher_assessmentsInterface {
    ID?: number;
    //Timestamp?: Date | null;
    Student?: StudentsInterface;
    Student_ID?: number;     // foreignkey.ID?
    Teacher?: TeachersInterface;
    Teacher_ID?: number;     // foreignkey.ID?
    Teaching_duration?: Teaching_durationsInterface;
    Teaching_duration_ID?: number;     // foreignkey.ID?
    Content_difficulty_level?: Content_difficulty_levelsInterface;
    Content_difficulty_level_ID?: number;     // foreignkey.ID?
    Content_quality?: Content_qualitysInterface;
    Content_quality_ID?: number;     // foreignkey.ID?
    Comment?: string;

    
}
