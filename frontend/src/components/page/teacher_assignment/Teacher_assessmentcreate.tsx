import React, { useEffect, useState } from "react";
import { Link as RouterLink } from "react-router-dom";
import Button from "@mui/material/Button";
import FormControl from "@mui/material/FormControl";
import Container from "@mui/material/Container";
import Paper from "@mui/material/Paper";
import Grid from "@mui/material/Grid";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import Divider from "@mui/material/Divider";
import Snackbar from "@mui/material/Snackbar";
import Select, { SelectChangeEvent } from "@mui/material/Select";
import MuiAlert, { AlertProps } from "@mui/material/Alert";
import TextField from "@mui/material/TextField";
// import { AdapterDateFns } from "@mui/x-date-pickers/AdapterDateFns";
// import { LocalizationProvider } from "@mui/x-date-pickers/LocalizationProvider";
// import { DatePicker } from "@mui/x-date-pickers/DatePicker";

 import { StudentsInterface } from "../../../models/IStudent";
 import { TeachersInterface } from "../../../models/ITeacher";
 import { Teaching_durationsInterface } from "../../../models/ITeaching_duration";
 import { Content_difficulty_levelsInterface } from "../../../models/IContent_difficulty_level";
 import { Content_qualitysInterface } from "../../../models/IContent_quality";
 import { Teacher_assessmentsInterface } from "../../../models/ITeacher_assessment";
 import { CommentsInterface } from "../../../models/IComment";     

import {
    GetComment,
    GetOnlyStudent,
    GetTeacher,
    GetTeaching_duration,
    GetContent_difficulty_level,
    GetContent_quality,
    CreateTeacher_assessment,
  } from "../../../services/HttpClientService";
  const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
    props,
    ref
  ) {
    return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
  });
  
  function Teacher_assessmentCreate() {
    const [student, setStudent] = useState<StudentsInterface>({});
    const [Teachers, setTeachers] = useState<TeachersInterface[]>([]);
    const [Teaching_durations, setTeaching_durations] = useState<Teaching_durationsInterface[]>([]);
    const [Content_difficulty_levels, setContent_difficulty_levels] = useState<Content_difficulty_levelsInterface[]>([]);
    const [Content_qualitys, setContent_qualitys] = useState<Content_qualitysInterface[]>([]);
    const [Teacher_assessment, setTeacher_assessments] = useState<Teacher_assessmentsInterface>({});
    const [comment, setcomments] = useState<string>("");

    const [success, setSuccess] = useState(false);
    const [error, setError] = useState(false);

    const handleClose = (
      event?: React.SyntheticEvent | Event,
      reason?: string
    ) => {
      if (reason === "clickaway") {
        return;
      }
      setSuccess(false);
      setError(false);
    };

    const handleChange = (event: SelectChangeEvent) => {
      const name = event.target.name as keyof typeof Teacher_assessment;
      const value = event.target.value;
      setTeacher_assessments({
          ...Teacher_assessment,
          [name]: value,
      });
      console.log(`${name}: ${value}`);
  };
    const getStudent = async () => {
      let res = await GetOnlyStudent();
      if (res) {
        setStudent(res);
        Teacher_assessment.Student_ID = res.ID
        console.log(res);
    }
  };
    const getcomment = async () => {
     let res = await GetComment();
      if (res) {
       setcomments(res);
       console.log(res);
  }
}; 
    const getTeacher = async () => {
      let res = await GetTeacher();
        if (res) {
          setTeachers(res);
          console.log(res);
  }
};

    const getTeaching_duration = async () => {
      let res = await GetTeaching_duration();
      if (res) {
        setTeaching_durations(res);
        console.log(res);
    }
  };
      const getContent_difficulty_level = async () => {
      let res = await GetContent_difficulty_level();
      if (res) {
        setContent_difficulty_levels(res);
        console.log(res);
    }
  };
      const getContent_quality = async () => {
      let res = await GetContent_quality();
      if (res) {
        setContent_qualitys(res);
        console.log(res);
    }
  };

  useEffect(() => {
    getcomment();
    getStudent();
    getTeacher();
    getTeaching_duration();
    getContent_difficulty_level();
    getContent_quality();

  }, []);

  const convertType = (data: string | number | undefined) => {
    let val = typeof data === "string" ? parseInt(data) : data;
    return val;
  };

  async function submit() {
    let data = {
      Student_ID: convertType(Teacher_assessment.Student_ID),
      Teacher_ID: convertType(Teacher_assessment.Teacher_ID),
      Teaching_duration_ID: convertType(Teacher_assessment.Teaching_duration_ID),
      Content_difficulty_level_ID: convertType(Teacher_assessment.Content_difficulty_level_ID),
      Content_quality_ID: convertType(Teacher_assessment.Content_quality_ID),
      Comment: (comment),

    };
    console.log("data");
    console.log(data);
    let res = await CreateTeacher_assessment(data);
    if (res) {
      setSuccess(true);
    } else {
      setError(true);
    }
  }
  return (
    <Container maxWidth="md">
      <Snackbar
        open={success}
        autoHideDuration={3000}
        onClose={handleClose}
        anchorOrigin={{ vertical: "top", horizontal: "center" }}
      >
        <Alert onClose={handleClose} severity="success">
          บันทึกข้อมูลสำเร็จ
        </Alert>
      </Snackbar>
      <Snackbar
        open={error}
        autoHideDuration={6000}
        onClose={handleClose}
        anchorOrigin={{ vertical: "top", horizontal: "center" }}
      >
        <Alert onClose={handleClose} severity="error">
          บันทึกข้อมูลไม่สำเร็จ
        </Alert>
      </Snackbar>
      <Paper>
        <Box
          display="flex"
          sx={{
            marginTop: 2,
          }}
        >
          <Box sx={{ paddingX: 2, paddingY: 1 }}>
            <Typography
              component="h2"
              variant="h6"
              color="primary"
              gutterBottom
            >
              ประเมินผู้สอน
            </Typography>
          </Box>
        </Box>
        <Divider />
        <Grid container spacing={3} sx={{ padding: 2 }}>
        <Grid item xs={8}>
            <FormControl fullWidth variant="outlined">
              <p>รหัสนักศึกษา</p>
              <Select
                native
                
                value={Teacher_assessment.Student_ID + ""}
                onChange={handleChange}
                disabled
                inputProps={{
                  name: "Student_ID",
                }}
              >
                <option aria-label="None" value="">
                  กรุณารหัสนักศึกษา
                </option>
                <option value={student?.ID} key={student?.ID}>
                  {student?.S_ID}
                </option>
              </Select>
            </FormControl>
          </Grid>
          <Grid item xs={8}>
            <FormControl fullWidth variant="outlined">
              <p>อาจารย์ผู้สอน</p>
              <Select
                native
                value={Teacher_assessment.Teacher_ID + ""}
                onChange={handleChange}
                inputProps={{
                  name: "Teacher_ID",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือกอาจารย์ผู้สอน
                </option>
                {Teachers.map((item: TeachersInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Name}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>
          <Grid item xs={8}>
            <FormControl fullWidth variant="outlined">
              <p>ระยะเวลาในการสอน</p>
              <Select
                native
                value={Teacher_assessment.Teaching_duration_ID + ""}
                onChange={handleChange}
                inputProps={{
                  name: "Teaching_duration_ID",
                }}
              >
                <option aria-label="None" value="">
                  เลือกระดับประเมิน
                </option>
                {Teaching_durations.map((item: Teaching_durationsInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Description}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>
          <Grid item xs={8}>
            <FormControl fullWidth variant="outlined">
              <p>ระดับความยากของเนื้อหา</p>
              <Select
                native
                value={Teacher_assessment.Content_difficulty_level_ID + ""}
                onChange={handleChange}
                inputProps={{
                  name: "Content_difficulty_level_ID",
                }}
              >
                <option aria-label="None" value="">
                เลือกระดับประเมิน
                </option>
                {Content_difficulty_levels.map((item: Content_difficulty_levelsInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Description}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>
          <Grid item xs={8}>
            <FormControl fullWidth variant="outlined">
              <p>คุณภาพของเนื้อหา</p>
              <Select
                native
                value={Teacher_assessment.Content_quality_ID + ""}
                onChange={handleChange}
                inputProps={{
                  name: "Content_quality_ID",
                }}
              >
                <option aria-label="None" value="">
                  เลือกระดับประเมิน
                </option>
                {Content_qualitys.map((item: Content_qualitysInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Description}
                  </option>
                  
                ))}
              </Select>
            </FormControl>
          </Grid>
          <Grid item xs={8}>
                <p>ความคิดเห็นเพิ่มเติม</p>
                <TextField fullWidth id="Comment" type="string" variant="outlined"  onChange={(event) => setcomments(event.target.value)} />
              </Grid>
          <Grid item xs={12}>
            <Button
              component={RouterLink}
              to="/Teacher_assessment"
              variant="contained"
              color="inherit"
            >
              กลับ
            </Button>
            <Button
              style={{ float: "right" }}
              onClick={submit}
              variant="contained"
              color="primary"
            >
              บันทึก
            </Button>
          </Grid>
        </Grid>
      </Paper>
    </Container>
  );
}

export default Teacher_assessmentCreate;