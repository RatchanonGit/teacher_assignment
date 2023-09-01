import React, { useState, useEffect } from "react";
import { Link as RouterLink } from "react-router-dom";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import Container from "@mui/material/Container";
import Box from "@mui/material/Box";
import { DataGrid, GridColDef } from "@mui/x-data-grid";
import { Teacher_assessmentsInterface } from "../../../models/ITeacher_assessment";
import { GetTeacher_assessments } from "../../../services/HttpClientService";


function Teacher_assessmentList() {
    const [Teacher_assessments, setTeacher_assessments] = useState<Teacher_assessmentsInterface[]>([]);
  
    useEffect(() => {
      getTeacher_assessments();
    }, []);
  
    const getTeacher_assessments = async () => {
      let res = await GetTeacher_assessments();
      if (res) {
        console.log("res");
        console.log(res);
        setTeacher_assessments(res);
      } 
    };
  
    const columns: GridColDef[] = [
      { field: "ID", headerName: "ลำดับ", width: 50 },
      {
        field: "Student",
        headerName: "นักศึกษา",
        width: 300,
        valueFormatter: (params) => params.value.Name,
      },
      {
        field: "Teacher",
        headerName: "อาจารย์ผู้สอน",
        width: 300,
        valueFormatter: (params) => params.value.Name,
      },
      {
        field: "Teaching_duration",
        headerName: "ระยะเวลาที่สอน",
        width: 300,
        valueFormatter: (params) => params.value.Description,
      },
      {
        field: "Content_difficulty_level",
        headerName: "ระดับความยากของเนื้อหา",
        width: 300,
        valueFormatter: (params) => params.value.Description,
      },
      {
        field: "Content_quality",
        headerName: "คุณภาพของเนื้อหา",
        width: 300,
        valueFormatter: (params) => params.value.Description,
      },
      {
        field: "Comment",
        headerName: "ความคิดเห็น",
        width: 300,
      },
    ];
  
    return (
      <div>
        <Container maxWidth="md">
          <Box
            display="flex"
            sx={{
              marginTop: 2,
            }}
          >
            <Box flexGrow={1}>
              <Typography
                component="h2"
                variant="h6"
                color="primary"
                gutterBottom
              >
                ข้อมูลการประเมินผู้สอน
              </Typography>
            </Box>
            <Box>
              <Button
                component={RouterLink}
                to="/Teacher_assessment/create"
                variant="contained"
                color="primary"
              >
                เลือกประเมินผู้สอนท่านอื่น
              </Button>
            </Box>
          </Box>
          <div style={{ height: 400, width: "100%", marginTop: "20px" }}>
            <DataGrid
              rows={Teacher_assessments}
              getRowId={(row) => row.ID}
              columns={columns}
              pageSize={5}
              rowsPerPageOptions={[5]}
            />
          </div>
        </Container>
      </div>
    );
  }
  
  export default Teacher_assessmentList;