"use client";
import { Tabs, TabList, TabPanels, Tab, TabPanel, Flex, Center, Box, Grid } from "@chakra-ui/react";
import RankingTable from "../organism/RankingTable";
import { Pagination } from "../molecules/Pagination";
import { CustomButton } from "../atoms/CustomButton";
import RefreshButton from "../atoms/RefreshButton";
import { useEffect, useState } from "react";
// import { client } from "@/libs/api";

const RankingTabs = () => {
  const [scoreRankings, setScoreRankings] = useState<ScoreRanking[]>(demoAccuracyRankings);
  const [rankingStartFrom, setRankingStartFrom] = useState(0);
  const [sortBy, setSortBy] = useState<"accuracy" | "keystrokes">("accuracy");
  const LIMIT = 10;

  const MAXIMUM = 100; // TODO: MAXIMUMをAPIから取得する

  useEffect(() => {
    // ページが読み込まれたときにデータを取得
    fetchData();
  });

  const handleTabChange = (index: number) => {
    if (index === 0) {
      setSortBy("accuracy");
    } else if (index === 1) {
      setSortBy("keystrokes");
    }

    fetchData;
  };

  // 演算子を引数にとる、ボタンを押したときのハンドラ関数
  const handlePaginationClick = (direction: "next" | "prev") => {
    const newStartFrom =
      direction === "prev" ? Math.max(rankingStartFrom - LIMIT, 0) : Math.min(rankingStartFrom + LIMIT, MAXIMUM);
    setRankingStartFrom(newStartFrom);

    fetchData;
  };

  const fetchData = async () => {
    // TODO: APIを使ってデータをフェッチ
    if (sortBy == "accuracy") {
      setScoreRankings(demoAccuracyRankings);
    } else if (sortBy == "keystrokes") {
      setScoreRankings(demoKeyStrokeRankings);
    }
  };

  return (
    <Tabs onChange={handleTabChange}>
      <Flex justifyContent={"center"}>
        <Grid templateColumns={"repeat(3, 1fr)"} gap={"300px"}>
          <Box opacity={"0"}>{/* 幅を揃えるためだけの要素，視覚的な意味はなし */}</Box>
          <TabList color={"white"}>
            <Tab _selected={{ color: "#00ace6" }}>正打率</Tab>
            <Tab _selected={{ color: "#00ace6" }}>入力文字数</Tab>
          </TabList>
          <RefreshButton onClick={() => fetchData()} isDisabled={false}/>
        </Grid>
      </Flex>

      <TabPanels>
        <TabPanel>
          <RankingTable scoreRankings={scoreRankings} />
        </TabPanel>
        <TabPanel>
          <RankingTable scoreRankings={scoreRankings} />
        </TabPanel>
      </TabPanels>
      <Center>
        <Pagination
          onPrev={() => handlePaginationClick("prev")}
          onNext={() => handlePaginationClick("next")}
          isPrevDisabled={rankingStartFrom <= 0}
          isNextDisabled={rankingStartFrom + LIMIT >= MAXIMUM}
        />
      </Center>
    </Tabs>
  );
};
export default RankingTabs;

export interface User {
  id: string;
  studentNumber: string;
  handleName: string;
}

export interface ScoreRanking {
  rank: Number;
  user: User;
  keystrokes: Number;
  accuracy: Number;
  createdAt: Date;
}

const demoUsers: User[] = [
  {
    id: "1",
    studentNumber: "70310000",
    handleName: "X",
  },
  {
    id: "2",
    studentNumber: "70310000",
    handleName: "Y",
  },
  {
    id: "3",
    studentNumber: "70310000",
    handleName: "Z",
  },
];

const demoKeyStrokeRankings: ScoreRanking[] = [
  {
    rank: 1,
    user: demoUsers[0],
    keystrokes: 100,
    accuracy: 100,
    createdAt: new Date(),
  },
  {
    rank: 2,
    user: demoUsers[1],
    keystrokes: 90,
    accuracy: 90,
    createdAt: new Date(),
  },
  {
    rank: 3,
    user: demoUsers[2],
    keystrokes: 80,
    accuracy: 80,
    createdAt: new Date(),
  },
  {
    rank: 4,
    user: demoUsers[2],
    keystrokes: 70,
    accuracy: 70,
    createdAt: new Date(),
  },
];

const demoAccuracyRankings: ScoreRanking[] = [
  {
    rank: 1,
    user: demoUsers[0],
    keystrokes: 100,
    accuracy: 100,
    createdAt: new Date(),
  },
  {
    rank: 2,
    user: demoUsers[1],
    keystrokes: 90,
    accuracy: 90,
    createdAt: new Date(),
  },
  {
    rank: 3,
    user: demoUsers[2],
    keystrokes: 80,
    accuracy: 80,
    createdAt: new Date(),
  },
  {
    rank: 4,
    user: demoUsers[2],
    keystrokes: 70,
    accuracy: 70,
    createdAt: new Date(),
  },
  {
    rank: 5,
    user: demoUsers[2],
    keystrokes: 60,
    accuracy: 60,
    createdAt: new Date(),
  },
];
