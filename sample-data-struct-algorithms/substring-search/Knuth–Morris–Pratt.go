package main

func main() {

}

type pat string

func (this *pat) KMP(pattern string) {

	this.pat = pattern

}

// public KMP(String pat)
// {
//  this.pat = pat;
//  M = pat.length();
//  dfa = new int[R][M];
//  dfa[pat.charAt(0)][0] = 1;
//  for (int X = 0, j = 1; j < M; j++)
//  {
//  for (int c = 0; c < R; c++)
//  dfa[c][j] = dfa[c][X];
//  dfa[pat.charAt(j)][j] = j+1;
//  X = dfa[pat.charAt(j)][X];
